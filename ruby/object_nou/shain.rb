# coding: utf-8
require_relative 'gyoumu'
require_relative 'yakushoku'

# Shain class
class Shain
    attr_accessor :gyoumu
    attr_accessor :yakushoku

    def initialize
        @gyoumu = Gyoumu.new
        @yakushoku = Yakushoku.new
    end

    def calculate_salary(kihonkyu)
        @yakushoku.calculate_salary(kihonkyu)
    end

    def up
        @yakushoku = @yakushoku.up
    end

    def down
        @yakushoku = @yakushoku.down
    end

    def work
        @gyoumu.standup
    end
end

