package adventofcode.year2020;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day7 extends BaseTest {

	@Test @Ignore public void singleCheck() {
		final ArrayList<Bag> bagList = new ArrayList<Bag>();
		System.out.println(parseBag("light red bags contain 1 bright white bag, 2 muted yellow bags.", bagList));
		System.out.println(parseBag("bright white bags contain 1 shiny gold bag.", bagList));
		System.out.println(parseBag("dotted black bags contain no other bags.", bagList));
		System.out.println(bagList);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day7/input.txt");
		final ArrayList<Bag> bagList = new ArrayList<Bag>();
		for (final String input : data) {
			parseBag(input, bagList);
		}
		final Bag current = bagList.stream().filter(l -> l.getColour().equals("shiny gold")).findFirst().orElse(null);
		System.out.println(current.bagContains());
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day7/input.txt");
		int count = 0;
		final ArrayList<Bag> bagList = new ArrayList<Bag>();
		for (final String input : data) {
			parseBag(input, bagList);
		}
		for (final Bag bag : bagList) {
			if (bag.isBagContains("shiny gold")) {
				count++;
			}
		}
		System.out.println(count);
	}

	private Bag parseBag(final String input, final ArrayList<Bag> bagsList) {
		final String bagColour = input.substring(0, input.indexOf(" bags contain "));

		Bag current = bagsList.stream().filter(l -> l.getColour().equals(bagColour)).findFirst().orElse(null);
		if (current == null) {
			current = new Bag(bagColour);
			bagsList.add(current);
		}
		final String containsBags = input.substring(bagColour.length() + 14);
		final String[] containsBagsArr = StringUtils.split(containsBags, ",.");
		for (final String b : containsBagsArr) {
			final String innerBag = b.trim();
			if (innerBag.startsWith("no")) {
			} else {
				final String[] innerBagArray = StringUtils.split(innerBag, " ");
				final int innerBagCount = Integer.parseInt(innerBagArray[0]);
				final String innerBagColour = innerBagArray[1] + " " + innerBagArray[2];
				//				System.out.println(innerBagCount + " -> " + innerBagColour);
				Bag currentChild = bagsList.stream().filter(l -> l.getColour().equals(innerBagColour)).findFirst().orElse(null);
				if (currentChild == null) {
					currentChild = new Bag(innerBagColour);
					bagsList.add(currentChild);
				}
				current.addBag(currentChild, innerBagCount);
			}
		}
		return null;
	}

	@Data class Bag {

		String colour;
		List<Bag> containsBags;
		Map<Bag, Integer> bagsCountMap = new HashMap<>();

		public Bag(final String colour) {
			this.colour = colour;
			this.containsBags = new ArrayList<Bag>();
		}

		public void addBag(final Bag bag, final int count) {
			final Bag current = this.containsBags.stream().filter(l -> l.getColour().equals(bag.getColour())).findFirst().orElse(null);
			if (current == null) {
				//				current = new Bag(bagColour);
				this.containsBags.add(bag);
			}
			this.bagsCountMap.put(bag, count);
			//			count += count;
		}

		public boolean isBagContains(final String bagColour) {
			final Bag current = this.containsBags.stream().filter(l -> l.getColour().equals(bagColour)).findFirst().orElse(null);
			if (current != null) {
				return true;
			}
			for (final Bag innerBag : this.containsBags) {
				if (innerBag.isBagContains(bagColour)) {
					return true;
				}
			}
			return false;
		}

		public int bagContains() {
			int totalCount = 0;
			for (final Bag innerBag : this.containsBags) {
				totalCount += this.bagsCountMap.get(innerBag) * (innerBag.bagContains() + 1);
			}
			return totalCount;
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Bag bag = (Bag) o;

			return this.colour != null ? this.colour.equals(bag.colour) : bag.colour == null;
		}

		@Override public int hashCode() {
			return this.colour != null ? this.colour.hashCode() : 0;
		}

		@Override public String toString() {
			return "Bag{" + "colour='" + this.colour + '\'' + ", containsBags=" + this.containsBags + '}' + "\n";
		}
	}
}
