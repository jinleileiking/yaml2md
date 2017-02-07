# require 'rest-client'
require 'trollop'
require 'pp'
require 'json'
require 'yaml'
require 'liquid'
# require 'file'


def get_val(input_item)
	return unless input_item
	@constants.each do |item| 
		if input_item["name"] == item["name"]

            # keys = ["require", "comment", "type", "must"]
            keys = item.keys.reject{|key| key == "name"}

            keys.each do |key|
                next if input_item[key]
			    input_item[key] = item[key] if item[key]
            end
            # input_item.reject {|key| key == "name"}.each do |key, val|
            #     next if keys.include? key
			#     input_item[key] = item[key] if item[key]
            # end
			# input_item["require"] = item["require"] if item["require"]
			# input_item["comment"] = item["comment"] if item["comment"]
			# input_item["type"] = item["type"] if item["type"]
			# input_item["must"] = item["must"] if item["must"] 
		end
	end
end





opts = Trollop::options do
      opt :file, "yaml file", :type => :string
end


unless opts[:file]
    puts "needs file param"
    exit
end

# puts File.basename opts[:file], ".yaml" 
data = YAML.load_file(opts[:file])

@constants = YAML.load_file('constant.yaml')

data["action"] = File.basename opts[:file], ".yaml"


if data["input"]
	data["input"].each do |item| 
		get_val(item)
	end
end

if data["output"]
	data["output"].each do |item| 
		get_val(item)
	end
end

data["detail"] = data["brief"] unless data["detail"]


template_file  = File.read("template.liquid")
template = Liquid::Template.parse(template_file)
out = template.render('data' => data)
print out

