# frozen_string_literal: true

require 'test_helper'

class HealthCheckControllerTest < ActionDispatch::IntegrationTest
  test 'should get index' do
    get '/health-check'
    assert_response :success
  end
end
