import os


def fill_template(template_name, **kwargs):
    import jinja2

    env = jinja2.Environment(
        loader=jinja2.FileSystemLoader(os.path.join(os.path.dirname(__file__), "templates")),
        autoescape=jinja2.select_autoescape(),
        trim_blocks=True,
    )
    templ = env.get_template(template_name)
    return templ.render(**kwargs)
