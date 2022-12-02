class Cell
  def initialize(number)
    @number = number
    @marked = false
  end

  def mark(number)
    @marked = true if @number == number
  end
end

class Board
  def initialize(lines)
    @cells = lines.map { |numbers| numbers.map { |number| Cell.new(number) } }
  end

  def mark(number)
    @cells.each do |line|
      line.each do |cell|
        cell.mark(number)
      end
    end
  end
end

lines = File.new('input.txt').read.split("\n\n")
input_numbers = lines.pop.split(',').map(&:to_i)

boards = lines.map do |board_input|
  board_lines = board_input.split("\n").map { |board_line| board_line.split(',') }
  Board.new(board_lines)
end

input_numbers.each { |number| boards.each { |board| board.mark(number) } }
puts boards.inspect
