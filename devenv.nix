{ pkgs, ... }:

{
  languages.go = {
    enable = true;
    package = pkgs.go_1_22;
  };
}
