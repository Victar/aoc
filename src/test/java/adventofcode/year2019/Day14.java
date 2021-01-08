package adventofcode.year2019;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import adventofcode.year2020.Day7;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;

public class Day14  extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test
	public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day14/input_sample.txt");
		ArrayList<Reaction> reactions = new ArrayList<>();
		for (String input : data) {

			String rIngr = input.substring(0, input.indexOf(" => "));
			String rStr = input.substring(input.indexOf(" => ")+4);
			String[] rNameArr = StringUtils.split(rStr, " ");
			Float currentCount = Float.parseFloat(rNameArr[0]);
			Reaction current = reactions.stream().filter(l -> l.getName().equals(rNameArr[1])).findFirst().orElse(null);
			if (current == null) {
				current = new Reaction(rNameArr[1]);
				reactions.add(current);
			}
			String[] rIngrArr = StringUtils.split(rIngr, " ,");
			for (int i=0; i<rIngrArr.length; i=i+2) {
				Float count = Float.parseFloat(rIngrArr[i]);
				String ingrName = rIngrArr[i+1];
				Reaction ingrReaction = reactions.stream().filter(l -> l.getName().equals(ingrName)).findFirst().orElse(null);
				if (ingrReaction == null) {
					ingrReaction = new Reaction(ingrName);
				}
				current.addReaction(ingrReaction, count/currentCount);
			}
		}
		Reaction fuel = reactions.stream().filter(l -> l.getName().equals("C")).findFirst().orElse(null);

		System.out.println(fuel.getCount());
	}

	@Data
//	@EqualsAndHashCode(of = "name")
//    @AllArgsConstructor(staticName = "of")
	class Reaction{
		String name;
		Map<Reaction, Float> reactionNeeded= new HashMap<>();

		public void addReaction(Reaction reaction, Float count){
			reactionNeeded.put(reaction, count);
		}
		public Reaction(final String name) {
			this.name = name;
		}

		public Float getCount(){
			if ("ORE".equals(name)){
				return 1.f;
			}
			float total = 0.f;
			for (Map.Entry<Reaction, Float> entry : reactionNeeded.entrySet()) {
				total += entry.getValue() * entry.getKey().getCount();
			}
			return total;
		}
	}
}

