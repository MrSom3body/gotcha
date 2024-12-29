{
  buildGoModule,
  interface ? "wlp2s0",
}:
buildGoModule {
  pname = "gotcha";
  version = "0.1.1";
  src = ./..;
  vendorHash = "sha256-hocnLCzWN8srQcO3BMNkd2lt0m54Qe7sqAhUxVZlz1k=";

  env.CGO_ENABLED = 0;

  ldflags = [
    "-s -w"
    "-X 'github.com/MrSom3body/gotcha/cmd.ifaceName=${interface}'"
    "-extldflags '-static'"
  ];
}
