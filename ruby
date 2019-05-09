#!/usr/bin/env ruby

require "hilighter"
require "io/wait"
require "optparse"

### Helpers begin
def err(msg)
    puts("[!] #{msg}".red)
end
def errx(msg)
    raise Exception.new(msg)
end
def good(msg)
    puts("[+] #{msg}".green)
end
def info(msg)
    puts("[*] #{msg}".white)
end
def subinfo(msg)
    puts("[=] #{msg}".cyan)
end
def warn(msg)
    puts("[-] #{msg}".yellow)
end
### Helpers end

class Exit
    GOOD = 0
    INVALID_OPTION = 1
    INVALID_ARGUMENT = 2
    MISSING_ARGUMENT = 3
    EXTRA_ARGUMENTS = 4
    EXCEPTION = 5
    AMBIGUOUS_ARGUMENT = 6
end

def parse(args)
    options = Hash.new
    options["verbose"] = false

    info = "TODO"

    parser = OptionParser.new do |opts|
        opts.summary_width = 20

        opts.banner = "Usage: #{File.basename($0)} [OPTIONS]"

        opts.on("")

        info.scan(/\S.{0,80}\S(?=\s|$)|\S+/).each do |line|
            opts.on("#{line}")
        end

        opts.on("", "OPTIONS")

        opts.on("-h", "--help", "Display this help message") do
            puts(opts)
            exit Exit::GOOD
        end

        opts.on("--[no-]color", "Disable colorized output") do |c|
            Hilighter.disable if (c)
        end

        opts.on(
            "-s",
            "--store=VALUE",
            "Example for storing cli arg"
        ) do |value|
            options["value"] = value
        end

        opts.on(
            "-v",
            "--verbose",
            "Show backtrace when error occurs"
        ) do
            options["verbose"] = true
        end
    end

    begin
        parser.parse!(args)
    rescue OptionParser::InvalidOption => e
        puts(e.message)
        puts(parser)
        exit Exit::INVALID_OPTION
    rescue OptionParser::InvalidArgument => e
        puts(e.message)
        puts(parser)
        exit Exit::INVALID_ARGUMENT
    rescue OptionParser::MissingArgument => e
        puts(e.message)
        puts(parser)
        exit Exit::MISSING_ARGUMENT
    rescue OptionParser::AmbiguousOption => e
        puts(e.message)
        puts(parser)
        exit Exit::AMBIGUOUS_ARGUMENT
    end

    if (args.empty?)
        puts(parser)
        exit Exit::MISSING_ARGUMENT
    elsif (!args.empty?)
        puts(parser)
        exit Exit::EXTRA_ARGUMENTS
    end

    return options
end

begin
    options = parse(ARGV)
rescue Interrupt
    # Exit gracefully on ^C
    exit Exit::GOOD
end

begin
    # Do stuff here
rescue Interrupt
    # Exit gracefully on ^C
rescue Errno::EPIPE
    # Do nothing. This can happen if piping to another program such as
    # less. Usually if less is closed before we're done with STDOUT.
rescue Exception => e
    $stderr.puts(
        [
            "Oops! Looks like an error has occured! Maybe the",
            "message below will help. If not, you can use the",
            "--verbose flag to get a backtrace."
        ].join(" ").wrap
    )
    $stderr.puts

    $stderr.puts(e.message.white.on_red)
    if (options["verbose"])
        e.backtrace.each do |line|
            $stderr.puts(line.light_yellow)
        end
    end
    exit Exit::EXCEPTION
end
exit Exit::GOOD
