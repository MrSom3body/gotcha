{
  buildGoModule,
  interface ? "wlp2s0",
}:
buildGoModule rec {
  pname = "gotcha";
  version = "0.1.3";
  src = ./..;
  vendorHash = "sha256-hocnLCzWN8srQcO3BMNkd2lt0m54Qe7sqAhUxVZlz1k=";

  env.CGO_ENABLED = 0;

  ldflags = [
    "-s -w"
    "-X 'github.com/MrSom3body/gotcha/cmd.ifaceName=${interface}'"
    "-X 'github.com/MrSom3body/gotcha/cmd.version=${version}'"
    "-extldflags '-static'"
  ];
}
