let pkgs = import <nixpkgs> {}; in

pkgs.buildGoModule rec {
  pname = "lgtv-tbc";
  version = "dev";
  src = ./.;
  vendorSha256 = "0sjjj9z1dhilhpc8pq4154czrb79z9cm044jvn75kxcjv6v5l2m5";
}
