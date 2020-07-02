def _impl(ctx):
    ctx.actions.run(
        outputs = [ctx.outputs.out],
        arguments = [ctx.attr.name, "-o", ctx.outputs.out.path],
        progress_message = "Say hello to %s" % ctx.attr.name,
        executable = ctx.executable._hello,
    )

hello = rule(
    implementation = _impl,
    attrs = {
        "out": attr.output(mandatory = True),
        "_hello": attr.label(
            executable = True,
            cfg = "host",
            default = Label("//:bazel-action-run-test"),
        ),
    },
)
