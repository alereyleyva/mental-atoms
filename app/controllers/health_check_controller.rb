# frozen_string_literal: true

class HealthCheckController < ApplicationController
  def index
    render json: { message: 'mental atoms is up and running!!!' }
  end
end
