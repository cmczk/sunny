#!/bin/bash

if [ ! -d $HOME/.sunny ]; then
    mkdir -p $HOME/.sunny/lua
    touch $HOME/.sunny/.lua-version.global
fi

go build
mv sunny $HOME/.sunny

current_shell=$(basename $SHELL)

case "$current_shell" in
	bash)
		config_file="$HOME/.bashrc"
		;;
	zsh)
	    config_file="$HOME/.zshrc"
		;;
	fish)
	    config_file="$HOME/.config/fish/config.fish"
		;;
	*)
	    config_file=""
esac

if [ -n "$config_file" ]; then
    cat <<EOF >> $config_file
# sunny
SUNNY_PATH="\$HOME/.sunny"
if [ -d "\$HOME/.sunny" ]; then
    export PATH="\$SUNNY_PATH:\$PATH"
    eval \$(sunny path)
fi

EOF
else
    echo "cannot find shell config file"
fi
