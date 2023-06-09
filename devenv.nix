{ pkgs, ... }:

{
  # https://devenv.sh/languages/
  languages.go = {
    enable = true;
    package = pkgs.go_1_20;
  };
}
