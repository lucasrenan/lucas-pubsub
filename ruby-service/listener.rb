require 'google/cloud/pubsub'

pubsub = Google::Cloud::Pubsub.new(project: ENV['GCP_PROJECT_ID'])
subscription = pubsub.subscription(ENV['PUBSUB_SUBSCRIPTION'])

subscriber = subscription.listen do |received_message|
  puts "Received message: #{received_message.data}"

  unless received_message.attributes.empty?
    puts "Attributes:"
    received_message.attributes.each do |key, value|
      puts "  #{key}: #{value}"
    end
  end

  puts '--- ACK'
  received_message.acknowledge!
end

subscriber.start

loop {}
# subscriber.stop.wait!
