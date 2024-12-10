using System.Threading.Tasks;
using Khulnasoft;

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
