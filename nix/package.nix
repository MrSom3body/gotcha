{
  buildGoModule,
  interface ? "wlp2s0",
}:
buildGoModule {
  pname = "gotcha";
  version = "0.1.0";
  src = ./..;
  vendorHash = "sha256-hocnLCzWN8srQcO3BMNkd2lt0m54Qe7sqAhUxVZlz1k=";

  ldflags = [
    "-X 'github.com/MrSom3body/gotcha/cmd.IfaceName=${interface}'"
  ];
}
