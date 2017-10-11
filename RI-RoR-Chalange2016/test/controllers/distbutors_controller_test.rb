require 'test_helper'

class DistbutorsControllerTest < ActionDispatch::IntegrationTest
  setup do
    @distbutor = distbutors(:one)
  end

  test "should get index" do
    get distbutors_url
    assert_response :success
  end

  test "should get new" do
    get new_distbutor_url
    assert_response :success
  end

  test "should create distbutor" do
    assert_difference('Distbutor.count') do
      post distbutors_url, params: { distbutor: { name: @distbutor.name } }
    end

    assert_redirected_to distbutor_url(Distbutor.last)
  end

  test "should show distbutor" do
    get distbutor_url(@distbutor)
    assert_response :success
  end

  test "should get edit" do
    get edit_distbutor_url(@distbutor)
    assert_response :success
  end

  test "should update distbutor" do
    patch distbutor_url(@distbutor), params: { distbutor: { name: @distbutor.name } }
    assert_redirected_to distbutor_url(@distbutor)
  end

  test "should destroy distbutor" do
    assert_difference('Distbutor.count', -1) do
      delete distbutor_url(@distbutor)
    end

    assert_redirected_to distbutors_url
  end
end
