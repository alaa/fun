require 'set'

class Matrix
  def initialize(matrix)
    @matrix = matrix
    @moves = Set.new
    @previous = nil
  end

  def rotate
    swap(0, 0)
    swap(0, 1)
    swap(0, 2)
    swap(0, 3)

    p matrix[0]
    p matrix[1]
    p matrix[2]
    p matrix[3]
  end

  private

  attr_accessor :matrix, :moves

  def swap(row, index)
    unless record?(row, index)

      new_row = index
      new_index = (3 - row.to_i).abs

      if @previous.nil?
        @previous = matrix[new_row][new_index]
        matrix[new_row][new_index] = matrix[row][index]
        record(row, index)
      else
        temp = matrix[new_row][new_index]
        matrix[new_row][new_index] = @previous
        @previous = temp
        record(row, index)
      end

      swap(new_row, new_index) unless record?(new_row, new_index)
    end
    @previous = nil
  end

  def record(row, index)
    address = "#{row}:#{index}"
    moves.add(address)
  end

  def record?(row, index)
    address = "#{row}:#{index}"
    moves.include? address
  end
end

matrix = [
  ['A','B','C','D'],
  ['E','F','G','H'],
  ['I','J','K','L'],
  ['M','N','O','P']
]

m = Matrix.new(matrix)
m.rotate

# Expected Output
#
# ["M", "I", "E", "A"]
# ["N", "F", "G", "B"]
# ["O", "J", "K", "C"]
# ["P", "L", "H", "D"]

