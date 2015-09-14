class String
  def freq(limit = 1000000)
    word_list = self.split(//)
    count = Hash.new(0)
    for word in word_list do
      count[word] += 1
    end
    count.sort_by { |char, count| char }.first(limit)
  end
end

require 'rspec'

describe "freq" do
  it "Returns hash of first 3 char occurances" do
    str = "abcabcabcdef"
    expect(str.freq(3)).to eql [['a',3],['b',3],['c',3]]
  end
end
