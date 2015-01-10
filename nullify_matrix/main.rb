class Matrix

  attr_accessor :rows, :columns, :matrix

  def initialize(matrix)
    @matrix = matrix
    @rows = []
    @columns = []
  end

  def nullify
    scan_for_zeros
    matrix.each_with_index do |row, row_index|
      row.each_with_index do |cell, cell_index|
        if (rows.include?(row_index) || (columns.include?(cell_index)))
          matrix[row_index][cell_index] = 0
        end
      end
    end
  end

  def scan_for_zeros
    matrix.each_with_index do |row, row_index|
      row.each_with_index do |cell, cell_index|
        if (cell == 0)
          rows << row_index
          columns << cell_index
        end
      end
    end
  end
end

matrix = [
  [1,2,3,0],
  [4,5,6,7],
  [8,9,0,10],
  [1,1,3,4]
]

m = Matrix.new(matrix)
m.nullify

p m.matrix[0]
p m.matrix[1]
p m.matrix[2]
p m.matrix[3]

# expected output
#
# [0, 0, 0, 0]
# [4, 5, 0, 0]
# [0, 0, 0, 0]
# [1, 1, 0, 0]

