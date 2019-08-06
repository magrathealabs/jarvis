services = %w[
    jarvis_jarvis
    jarvis_rabbitmq
]

def now
    @now ||= Time.now
end

def actual_dir
    File.basename(Dir.getwd)
end

def git_project_prefix
    "#{actual_dir}_"
end

def dockerhub_repository
    "magrathealabs/jarvis"
end

def remove_prefix(service)
    service.gsub(git_project_prefix, '')
end

def dated
    [branch, now.year, now.month, now.day, now.hour, now.min].join('.')
end

def branch
    ENV['TRAVIS_BRANCH'] || 'unstable'
end

def tag(service)
    "#{remove_prefix(service)}-#{dated}".
        gsub('-', '.').
        gsub('/', '.').
        gsub('_', '.').
        gsub(' ', '.')
end

services.each do |service|
   system "docker tag #{service} #{dockerhub_repository}:#{tag(service)}"
   system "docker push #{dockerhub_repository}:#{tag(service)}"
end
