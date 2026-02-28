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
          ];

          shellHook = ''
            export CGO_ENABLED=1
            export CC=gcc
            export PKG_CONFIG_PATH=${pkgs.libX11}/lib/pkgconfig:${pkgs.mesa}/lib/pkgconfig:$PKG_CONFIG_PATH
            # Автоматически добавляет пути библиотек в LD_LIBRARY_PATH
            export LD_LIBRARY_PATH=${pkgs.libGL}/lib:${pkgs.mesa}/lib:${pkgs.libX11}/lib:$LD_LIBRARY_PATH
            echo "OpenGL libraries loaded from: ${pkgs.libGL}/lib"
            echo "Ebitengine development environment loaded"
          '';
        };
      });
}