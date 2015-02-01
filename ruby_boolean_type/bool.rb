module Boolean; end
class TrueClass; include Boolean; end
class FalseClass; include Boolean; end

puts true.is_a?(Boolean)
puts false.is_a?(Boolean)

