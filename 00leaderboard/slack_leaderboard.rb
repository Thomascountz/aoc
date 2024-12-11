require "httparty"
require "json"
require "time"
require "logger"
require "stringio"

LEADERBOARD_ID = ENV.fetch("LEADERBOARD_ID")
AOC_COOKIE = ENV.fetch("AOC_COOKIE")

LEADERBOARD_URL = "https://adventofcode.com/2024/leaderboard/private/view/#{LEADERBOARD_ID}"
CACHE_FILE = "./leaderboard_#{LEADERBOARD_ID}_cache.json"
CACHE_DURATION = 24 * 60 * 60 # 24 hours in seconds

LOGGER = Logger.new(StringIO.new)

def fetch_leaderboard_data
  fetch_cached_leaderboard || call_leaderboard_api
end

def fetch_cached_leaderboard
  # Check if cache file exists and is valid
  if File.exist?(CACHE_FILE)
    cached_data = JSON.parse(File.read(CACHE_FILE))
    cache_timestamp = Time.at(cached_data["cache_timestamp"])

    if Time.now.to_i - cache_timestamp.to_i < CACHE_DURATION # Within 24 hours
      LOGGER.info("Using cached leaderboard data from #{cache_timestamp}")
      cached_data["leaderboard"]
    else
      LOGGER.info("Cache expired.")
      nil
    end
  else
    LOGGER.info("Cache not found at #{CACHE_FILE}.")
    nil
  end
end

def call_leaderboard_api
  # Fetch new data if cache is expired or doesn't exist
  LOGGER.info("Fetching leaderboard data.")

  raise "ENV['AOC_COOKIE'] is not set." if AOC_COOKIE.nil?

  response = HTTParty.get("#{LEADERBOARD_URL}.json", headers: {"cookie" => AOC_COOKIE})

  if response.ok?
    JSON.parse(response.body).tap do |data|
      LOGGER.info("Updating cache.")
      File.write(CACHE_FILE, JSON.pretty_generate({
        "cache_timestamp" => Time.now.to_i,
        "leaderboard" => data
      }))
    end
  else
    LOGGER.error("Failed to fetch leaderboard data: #{response.code} #{response.message}")
    nil
  end
end

def generate_slack_payload(leaderboard)
  LOGGER.info("Generating Slack payload.")

  members = leaderboard["members"].values
  sorted_members = members.sort_by { |m| -m["local_score"] } # Sort by descending local_score

  blocks = [
    {
      type: "section",
      text: {
        type: "mrkdwn",
        text: ":christmas_tree: *Advent of Code #{leaderboard["event"]}*\n_Leaderboard: <#{LEADERBOARD_URL}|#{LEADERBOARD_ID}>_"
      }
    },
    {type: "divider"}
  ]

  sorted_members.each_with_index do |member, index|
    name = member["name"] || "Anonymous"
    stars = member["stars"]
    next if stars.zero? # Skip participants with zero stars

    local_score = member["local_score"]
    last_star_ts = member["last_star_ts"]
    last_star_time = (last_star_ts > 0) ? Time.at(last_star_ts).utc.strftime("%Y-%m-%d %H:%M:%S") : "N/A"

    rank_emoji = case index
    when 0 then ":trophy:" # First place
    when 1 then ":second_place_medal:" # Second place
    when 2 then ":third_place_medal:" # Third place
    else ":star:"
    end

    blocks << {
      type: "section",
      text: {
        type: "mrkdwn",
        text: "#{rank_emoji} *#{index + 1}. #{name}*\n*Stars*: #{stars}  |  *Score*: #{local_score}  |  *Last Star*: #{last_star_time}"
      }
    }

    # Add a divider for top performers (top 3)
    blocks << {type: "divider"} if index == 2
  end

  # Collect participants with zero stars
  zero_star_members = sorted_members.select { |m| m["stars"] == 0 }.map { |m| m["name"] || "Anonymous" }

  # Add a footer message if there are any zero-star participants
  unless zero_star_members.empty?
    blocks << {type: "divider"}
    blocks << {
      type: "context",
      elements: [
        {
          type: "mrkdwn",
          text: ":sleeping: Participants yet to earn a star: #{zero_star_members.join(", ")}"
        }
      ]
    }
  end

  blocks << {type: "divider"}

  blocks << {
    type: "context",
    elements: [
      {
        type: "mrkdwn",
        text: "<#{LEADERBOARD_URL}|View the leaderboard>"
      }
    ]
  }

  {blocks: blocks}
end

# Example usage
leaderboard = fetch_leaderboard_data
raise "Failed to fetch leaderboard data." unless leaderboard

slack_payload = generate_slack_payload(leaderboard)
puts JSON.pretty_generate(slack_payload)
