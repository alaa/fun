class Anagram
  def initialize(text)
    @text = text.split(/\s/)
  end

  def find
    @text.each_with_object(Hash.new []) do |word, acc|
      acc[word.chars.sort] += [word]
    end.values
  end
end

anagram = Anagram.new "who eats how seat sara aras"
p anagram.find
