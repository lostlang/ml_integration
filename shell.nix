{ pkgs ? import <nixpkgs> { }, }:

pkgs.mkShell {
  buildInputs = with pkgs; [ go go-swag ];

  shellHook = "";
}
