{
  description = "Ebitengine development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            gcc13
            binutils
            pkg-config
            go

            # Графические библиотеки
            libX11
            libXrandr
            libXi
            libXcursor
            libXinerama
            libXxf86vm
            libxcb
            libXext
            libXfixes
            libXrender
            mesa
            libGL
            libGLU
            glfw
            glew
            freeglut
            patchelf

            alsa-lib
            alsa-plugins
            pipewire
            pulseaudio
          ];

          shellHook = ''
            export CGO_ENABLED=1
            export CC=gcc
            export PKG_CONFIG_PATH=${pkgs.libX11}/lib/pkgconfig:${pkgs.mesa}/lib/pkgconfig:$PKG_CONFIG_PATH
            # Автоматически добавляет пути библиотек в LD_LIBRARY_PATH
            export LD_LIBRARY_PATH=${pkgs.libGL}/lib:${pkgs.mesa}/lib:${pkgs.libX11}/lib:$LD_LIBRARY_PATH
            echo "OpenGL libraries loaded from: ${pkgs.libGL}/lib"

            # Используем ALSA-плагины из того же pipewire
                        export ALSA_PLUGIN_DIR=${pkgs.pipewire}/lib/alsa-lib

                        # Создаем конфиг для ALSA через PipeWire
                        cat > ~/.asoundrc << EOF
                        pcm.!default {
                            type pipewire
                            hint.description "Default ALSA device via PipeWire"
                        }
                        ctl.!default {
                            type pipewire
                            hint.description "Default control via PipeWire"
                        }
            EOF

                        # Пробрасываем сокет PipeWire из системы
                        export XDG_RUNTIME_DIR=/run/user/$(id -u)
                        export PIPEWIRE_RUNTIME_DIR=$XDG_RUNTIME_DIR
                        export LD_PRELOAD=/nix/store/97wpr90s0krpq4rv9jdqbr8aavfifgd4-alsa-lib-1.2.15.1/lib/libasound.so.2

                        echo "Sound configured with PipeWire from ${pkgs.pipewire}"

            echo "Ebitengine development environment loaded"
          '';
        };
      });
}