{
  buildGoModule,
  color ? "6",
}:
buildGoModule rec {
  pname = "gotcha";
  version = "2.1.0";
  src = ./..;
  vendorHash = "sha256-hocnLCzWN8srQcO3BMNkd2lt0m54Qe7sqAhUxVZlz1k=";

  env.CGO_ENABLED = 0;

  ldflags = [
    "-s -w"
    "-X 'github.com/MrSom3body/gotcha/cmd.color=${color}'"
    "-X 'github.com/MrSom3body/gotcha/cmd.version=${version}'"
    "-extldflags '-static'"
  ];
}
