{ pkgs ? import <nixpkgs> { }, }:

pkgs.mkShell {
  buildInputs = with pkgs; [ go golangci-lint go-swag q-text-as-data ];

  shellHook = "";
}
