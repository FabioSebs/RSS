from feedgen.feed import FeedGenerator

# RSS GEN CLASS
class RSSGenerator:
    def __init__(self) -> None:
        self.site = []
        pass

    def add_sites(self, *args):
        for arg in args:
            self.site.append(arg)

    def gen_rss(self):
        fg = FeedGenerator()
        fg.id('http://lernfunk.de/media/654321')
        fg.title('Some Testfeed LOOOOOll')
        fg.author( {'name':'John Doe','email':'john@example.de'} )
        fg.link( href='http://example.com', rel='alternate' )
        fg.logo('http://ex.com/logo.jpg')
        fg.subtitle('This is a cool feed!')
        fg.link( href='http://larskiesow.de/test.atom', rel='self' )
        fg.language('en')
        atomfeed = fg.atom_str(pretty=True) # Get the ATOM feed as string
        fg.atom_file('atom.xml') # Write the ATOM feed to a file
        fg.rss_file('rss.xml') # Write the RSS feed to a file
        fe = fg.add_entry()
        fe.id('http://lernfunk.de/media/654321/1')
        fe.title('The First Episode')
        fe.link(href="http://lernfunk.de/feed")
        print(fg.rss_str(pretty=True))

