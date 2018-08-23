class Login
  def self.authenticate_user(user)
    puts "\nLogin with default Admin credentials!"
    print "Enter username: "
    username = gets.chomp
    print "Enter password: "
    password = gets.chomp

    @login = true
    while @login
      if username.eql?(user.username) && password.eql?(user.password)
        puts "\nLogged in successfully. Welcome #{user.username.capitalize}!"
        @login = false
      else
        puts "Invalid username or password. Please try again."
        self.authenticate_user(user)
      end
    end
  end

  def self.logout
    print "\nTo Logged out type 'logout': "
    logout = gets.chomp
    puts "\nLogged out successfully!" if logout == 'logout'
  end
end