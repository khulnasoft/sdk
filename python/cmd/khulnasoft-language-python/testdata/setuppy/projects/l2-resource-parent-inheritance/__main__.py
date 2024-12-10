import khulnasoft
import khulnasoft_simple as simple

provider = simple.Provider("provider")
parent1 = simple.Resource("parent1", value=True,
opts = khulnasoft.ResourceOptions(provider=provider))
child1 = simple.Resource("child1", value=True,
opts = khulnasoft.ResourceOptions(parent=parent1))
orphan1 = simple.Resource("orphan1", value=True)
parent2 = simple.Resource("parent2", value=True,
opts = khulnasoft.ResourceOptions(protect=True))
child2 = simple.Resource("child2", value=True,
opts = khulnasoft.ResourceOptions(parent=parent2))
orphan2 = simple.Resource("orphan2", value=True)
