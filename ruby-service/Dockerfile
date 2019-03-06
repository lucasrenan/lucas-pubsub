FROM ruby:2.6.1-slim-stretch

WORKDIR /app
COPY Gemfile* ./

RUN bundle install

COPY . ./

CMD ["ruby", "listener.rb"]
