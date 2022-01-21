{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell { buildInputs = [ minikube kubectl kubernetes-helm ]; }
