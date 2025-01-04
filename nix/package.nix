{
  buildGoModule,
  ifaceName ? "lo",
}:
buildGoModule rec {
  pname = "gotcha";
  version = "1.0.2";
  src = ./..;
  vendorHash = "sha256-hocnLCzWN8srQcO3BMNkd2lt0m54Qe7sqAhUxVZlz1k=";

  env.CGO_ENABLED = 0;

  ldflags = [
    "-s -w"
    "-X 'github.com/MrSom3body/gotcha/cmd.ifaceName=${ifaceName}'"
    "-X 'github.com/MrSom3body/gotcha/cmd.version=${version}'"
    "-extldflags '-static'"
  ];
}
