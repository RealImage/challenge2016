class HomesController< ApplicationController
	def index

	end

	def new_distributor
		@new_distributor = {"name": ""}
	end

	def create_distributor
		binding.pry
	end

end