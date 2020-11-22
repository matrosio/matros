Vagrant.configure("2") do |config|
    config.vm.box = "centos/8"
    config.vm.box_version = "1905.1"
    config.vm.network :private_network, ip: "192.168.0.42"

    config.vm.define "single" do |instance|
        instance.vm.hostname = "matros.single"
    end

    config.vm.provider :virtualbox do |vb|
        vb.customize [
        "modifyvm", :id,
        "--cpuexecutioncap", "50",
        "--memory", "256",
        ]
    end
end
