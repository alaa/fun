class Unique
  attr_reader :uniq, :words

  def initialize(text)
    @uniq = Hash.new
    @words = text.split(' ')
  end

  def parse
    words.each { |w| add_to_stack(w) }
    uniq
  end

  private

  def add_to_stack(word)
    uniq.keys.include?(word) ? uniq[word] += 1 : uniq[word] = 1
  end
end

describe Unique do
  let(:text) { "a b c d d e f" }
  let(:obj) { described_class.new(text) }

  it "initializes with the right values" do
    expect(obj.uniq).to eql Hash.new
    expect(obj.words).to eql %w(a b c d d e f)
  end

  it "returns the unique values" do
    hash = {'a'=>1, 'b'=>1, 'c'=>1, 'd'=>2, 'e'=>1, 'f'=>1}
    expect(obj.parse).to eql hash
  end
end

