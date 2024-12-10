import khulnasoft
import khulnasoft_secret as secret

res = secret.Resource("res",
    private="closed",
    public="open",
    private_data=secret.DataArgs(
        private="closed",
        public="open",
    ),
    public_data=secret.DataArgs(
        private="closed",
        public="open",
    ))
