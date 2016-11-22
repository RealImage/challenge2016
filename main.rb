require_relative "models/repo"

repo = Repo.new("./inputs/cities.csv", "./inputs/distributors.txt")

test_cases_count = 0
failures = 0
test_cases = File.open("./inputs/test_cases.txt").read
test_cases.each_line do |test_case|
  next if test_case.length == 0
  next if test_case =~ /^#/

  test_cases_count += 1

  distributor, region, expected = test_case.split(",").map(&:strip).map(&:upcase)
  actual = repo.can_distribute_in?(distributor, region)
  if expected.to_s != actual.to_s.upcase
    puts [distributor, ",",  region, " : ", " [EXPECTED] ", expected.downcase, " [ACTUAL] ", actual].join
    failures += 1
  end
end
puts "%d Test Cases. %d Passed. %d Failed." % [test_cases_count, test_cases_count - failures, failures]