{ pkgs ? import <nixpkgs> { } }:

let
in pkgs.mkShell {
  buildInputs = [
    pkgs.just # Just runner
    pkgs.go # Go compiler
    pkgs.reflex # Hot reloading
  ];

  # Shell hook to set up environment
  shellHook = "";
}
