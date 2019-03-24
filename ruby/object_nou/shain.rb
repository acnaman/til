# coding: utf-8
require_relative 'gyoumu'

# Shain class
class Shain
    attr_accessor :gyoumu

    def initialize
        @gyoumu = Gyoumu.new
    end

    def work
        @gyoumu.standup
    end
end

