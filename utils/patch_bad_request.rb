#!/usr/bin/env ruby

require 'json'
require 'logger'
require 'optparse'

options = { log_level: Logger::Severity::INFO }

OptionParser.new do |opts|
  opts.banner = "Usage: #{File.basename(__FILE__)} [options]"

  opts.on('-h', '--help', 'Prints this help') do
    puts opts
    exit
  end

  opts.on('-f FILE', '--file FILE', 'Swagger spec patch') { |f| options[:file] = File.expand_path(f) }
end.parse!

@logger = Logger.new(STDOUT, level: options[:log_level], progname: File.basename(__FILE__))

unless options[:file]
  @logger.error('A spec file to patch is required!')
  exit 1
end

@logger.info("Patching spec file #{options.fetch(:file)}")
spec_file = JSON.load_file!(options.fetch(:file))

VERBS = %w(get delete patch post put).freeze
ADDITIONAL_RESPONSES = {
  '400' => {
    'description' => 'Bad Request',
    'schema' => {
      'type' => 'object',
      'additionalProperties' => true,
    },
  },
  'default' => {
    'description' => 'Unexpected Response',
    'schema' => {
      'type' => 'object',
      'additionalProperties' => true,
    },
  },
}.freeze

patch_count = 0

spec_file.fetch('paths').each do |path, spec|
  @logger.debug("Path: #{path}")
  VERBS.each do |verb|
    responses = spec.dig(verb, 'responses')
    next if responses.nil? || ADDITIONAL_RESPONSES.keys.all? { |r| responses.key?(r) }

    @logger.info("Patching: #{path} | Verb: #{verb}")
    responses.merge!(ADDITIONAL_RESPONSES)
    patch_count += 1
  end
end

@logger.info("Patching complete, patched #{patch_count} path verbs") unless patch_count.eql?(0)

File.write(options.fetch(:file), JSON.pretty_generate(spec_file))
@logger.info('Done')
