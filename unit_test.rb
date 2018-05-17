require './csv_import_util'
require './distributor'
require './country'
require 'test/unit'

class TestCsvImportUtil < Test::Unit::TestCase
  def test_import
    CsvImportUtil.parse_cities_csv('./cities.csv')
    CsvImportUtil.parse_distributors_csv('./distributors.csv')
  end
end

class TestDistributor < Test::Unit::TestCase
  def test_validate_location
    distributor = Distributor.new({"Distributor Name" => "Distributor3",
      "Allowed locations"=>"INDIA","Unallowed locations" => "KERALA-INDIA",
      "Parent Name" => "Distributor1,Distributor2"})
    assert_equal(true, distributor.validate_location('allowed_locations',
      distributor.parse_location('INDIA')))
    assert_equal(false, distributor.validate_location('allowed_locations',
      distributor.parse_location('CANADA')))
  end

  def test_validate_input_location
    distributor = Distributor.new({"Distributor Name" => "Distributor4",
      "Allowed locations"=>"INDIA","Unallowed locations" => "",
      "Parent Name" => ""})
    assert_equal(["thanchavoor is not valid in the provided country. So answer is NO"],
      distributor.validate_input_location({country: 'india', state: 'thanchavoor', city: ''}))
  end

  def test_validate_permission
    distributor4 = Distributor.new({"Distributor Name" => "Distributor4",
      "Allowed locations"=> "INDIA, UNITEDSTATES",
      "Unallowed locations" => "KARNATAKA-INDIA, CHENNAI-TAMIL NADU-INDIA", "Parent name" => ""})
    distributor5 = Distributor.new({"Distributor Name" => "Distributor5",
      "Allowed locations"=>"INDIA", "Unallowed locations" => "TAMIL NADU-INDIA",
      "Parent name" => "Distributor4"})
    assert_equal(["Yes", []], distributor4.validate_permission("TAMIL NADU-INDIA"))
    assert_equal(["No", []], distributor5.validate_permission("TAMIL NADU-INDIA"))
    assert_equal(["No", ["tamilnadu is not valid in the provided country. So answer is NO"]],
      distributor5.validate_permission("TAMILNADU-INDIA"))
  end
end
