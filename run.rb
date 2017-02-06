Dir.glob('*.yaml').each do |file|
	next unless file[0] == file[0].upcase
	puts file
	system("ruby genmd.rb -f #{file} >| #{File.basename file, ".yaml"}.md")
end

