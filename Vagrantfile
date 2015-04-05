# -*- mode: ruby -*-
# vi: set ft=ruby :
VAGRANTFILE_API_VERSION = "2"
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.provision :shell, :path => "bootstrap.sh"
  config.vm.hostname = "arsdev"
  config.vm.post_up_message = "ARS - Airline Reservation System Development Install"
  config.vm.network :private_network, ip: "192.168.56.101"
  config.vm.network "forwarded_port", guest: 3500, host: 3500
  config.vm.synced_folder ".", "/ARS", id: "vagrant-root",
    owner: "vagrant",
    group: "vagrant",
    create: true
end
