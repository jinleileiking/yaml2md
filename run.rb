Dir.glob('*.yaml').each do |file|
	next unless file[0] == file[0].upcase
    next if file[0] == '_'
	puts file
	system("ruby gen.rb -f #{file} >| mds/#{File.basename file, ".yaml"}.md")
end

