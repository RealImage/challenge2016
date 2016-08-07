require_relative "distributor"
require "test/unit"
 
class TestDistributor < Test::Unit::TestCase
  def test_check_country_permission
    distributor = Distributor.new({"Distributor Name" => "Distributor1","Allowed locations"=>"INDIA,UNITEDSTATES","Unallowed locations" => "","Extends" => ""})
    assert_equal(true, distributor.eval_location("allowed_locations", distributor.parse_location("INDIA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("CHINA")))
  end

  def test_check_state_permission
    distributor = Distributor.new({"Distributor Name" => "Distributor1","Allowed locations"=>"KARNATAKA-INDIA","Unallowed locations" => "","Extends" => ""})
    assert_equal(true, distributor.eval_location("allowed_locations", distributor.parse_location("KARNATAKA-INDIA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("TAMILNADU-INDIA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("KARNATAKA-CHINA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("TAMILNADU-CHINA")))

  end

  def test_check_city_permission
    distributor = Distributor.new({"Distributor Name" => "Distributor1","Allowed locations"=>"CHENNAI-TAMILNADU-INDIA","Unallowed locations" => "","Extends" => ""})
    assert_equal(true, distributor.eval_location("allowed_locations", distributor.parse_location("CHENNAI-TAMILNADU-INDIA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("TAMILNADU-CHINA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("TAMILNADU-INDIA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("ERODE-TAMILNADU-INDIA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("BANGALORE-KARNATAKA-INDIA")))

    distributor = Distributor.new({"Distributor Name" => "Distributor1","Allowed locations"=>"KARNATAKA-INDIA","Unallowed locations" => "","Extends" => ""})
    assert_equal(true, distributor.eval_location("allowed_locations", distributor.parse_location("KARNATAKA-INDIA")))
    assert_equal(true, distributor.eval_location("allowed_locations", distributor.parse_location("BANGALORE-KARNATAKA-INDIA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("KARNATAKA-CHINA")))
    assert_equal(false, distributor.eval_location("allowed_locations", distributor.parse_location("BANGALORE-KARNATAKA-CHINA")))
  end

  def test_permission
    distributor1 = Distributor.new({"Distributor Name" => "Distributor1","Allowed locations"=>"INDIA,UNITEDSTATES","Unallowed locations" => "KARNATAKA-INDIA,CHENNAI-TAMILNADU-INDIA","Extends" => ""})
    distributor2 = Distributor.new({"Distributor Name" => "Distributor2","Allowed locations"=>"INDIA","Unallowed locations" => "TAMILNADU-INDIA","Extends" => "Distributor1"})
    assert_equal("No", distributor1.has_permission?("KARNATAKA-INDIA"))
    assert_equal("No", distributor2.has_permission?("KARNATAKA-INDIA"))

    assert_equal("Yes", distributor1.has_permission?("INDIA"))
    assert_equal("Yes", distributor2.has_permission?("INDIA"))

    assert_equal("Yes", distributor1.has_permission?("UNITEDSTATES"))
    assert_equal("No", distributor2.has_permission?("UNITEDSTATES"))
    assert_equal("No", distributor2.has_permission?("ILLINOIS-UNITEDSTATES"))
    
    assert_equal("Yes", distributor1.has_permission?("TAMILNADU-INDIA"))
    assert_equal("No", distributor2.has_permission?("TAMILNADU-INDIA"))
    assert_equal("No", distributor1.has_permission?("CHENNAI-TAMILNADU-INDIA"))
    assert_equal("No", distributor2.has_permission?("CHENNAI-TAMILNADU-INDIA"))
    assert_equal("Yes", distributor1.has_permission?("ERODE-TAMILNADU-INDIA"))
    assert_equal("No", distributor2.has_permission?("ERODE-TAMILNADU-INDIA"))
    assert_equal("Yes", distributor2.has_permission?("KERALA-INDIA"))
  end
end
