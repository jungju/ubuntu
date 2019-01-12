Vagrant.configure("2") do |config|
  config.vm.box = "jungju/ubuntu180401"
    config.vm.provider "virtualbox" do |vb|
      # Display the VirtualBox GUI when booting the machine
      vb.gui = false
   
      # Customize the amount of memory on the VM:
      vb.memory = "1024"
    end

    config.vm.provision "shell", inline: <<-SHELL
      apt-get update
    SHELL
end
