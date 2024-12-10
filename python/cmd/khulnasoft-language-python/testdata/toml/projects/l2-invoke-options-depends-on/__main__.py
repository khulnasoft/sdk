import khulnasoft
import khulnasoft_simple_invoke as simple_invoke

explicit_provider = simple_invoke.Provider("explicitProvider")
first = simple_invoke.StringResource("first", text="first hello")
data = simple_invoke.my_invoke_output(value="hello", opts=khulnasoft.InvokeOutputOptions(depends_on=[first]))
second = simple_invoke.StringResource("second", text=data.result)
khulnasoft.export("hello", data.result)
