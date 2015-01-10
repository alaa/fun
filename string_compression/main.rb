class Compress
  def initialize(text)
    @text = text
    @buffer = []
  end

  def compress
    text_array = text.split('')
    text_array.each_with_index do |char, index|
      if ((index == 0) || (char != text_array[index -1]))
        buffer << [char, 1]
      else
        buffer.last[1] += 1
      end
    end

    output
  end

  private

  attr_reader :text, :buffer

  def output
    compressed_string = buffer.flatten.join
    compressed_string.size > text.size ? text : compressed_string
  end
end

describe Compress do
  describe "#compress" do

    context "compressed string length less than original text" do
      subject { Compress.new("aaaabcccc") }

      it "compresses the string" do
        expect(subject.compress).to eq "a4b1c4"
      end
    end

    context "compressed string length greater than original text" do
      let(:text) { 'aaabccddeef' }
      subject { Compress.new(text) }

      it "returns the original string" do
        expect(subject.compress).to eq text
      end
    end
  end
end
