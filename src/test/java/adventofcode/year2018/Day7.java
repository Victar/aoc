package adventofcode.year2018;

import adventofcode.BaseTest;
import org.junit.Test;

import java.util.*;

public class Day7 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2018/day7/input.txt");
		Set<Character> steps = new HashSet<>();
		Map<Character, List<Character>> childParentMaps = new HashMap<>();
		for (final String input : data) {
			final Character first = input.charAt(5);
			final Character second = input.charAt(36);
			if (childParentMaps.containsKey(second)) {
				childParentMaps.get(second).add(first);
			} else {
				final List<Character> currList = new ArrayList<>();
				currList.add(first);
				childParentMaps.put(second, currList);
			}
			steps.add(first);
			steps.add(second);
		}

		final List<Character> finalOrder = new ArrayList<>();

		while (childParentMaps.size() > 0) {
			final List<Character> tmp = new ArrayList<>(steps);
			tmp.removeAll(childParentMaps.keySet());
			Collections.sort(tmp);
			if (tmp.size() > 0) {
				Character toRemove = tmp.get(0);
				finalOrder.add(toRemove);
				final Map<Character, List<Character>> newChildParentMaps = new HashMap<>();
				for (final Map.Entry<Character, List<Character>> entry : childParentMaps.entrySet()) {
					final List<Character> newList = new ArrayList<>(entry.getValue());
					newList.remove(toRemove);
					if (newList.size() > 0) {
						newChildParentMaps.put(entry.getKey(), newList);
					}
				}
				childParentMaps = newChildParentMaps;
				steps.remove(toRemove);
			}
		}
		if (steps.size() == 1) {
			finalOrder.addAll(steps);
		}
		finalOrder.forEach(System.out::print);

	}


	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2018/day7/sample.txt");
		Set<Character> steps = new HashSet<>();
		Map<Character, List<Character>> childParentMaps = new HashMap<>();
		for (final String input : data) {
			final Character first = input.charAt(5);
			final Character second = input.charAt(36);
			if (childParentMaps.containsKey(second)) {
				childParentMaps.get(second).add(first);
			} else {
				final List<Character> currList = new ArrayList<>();
				currList.add(first);
				childParentMaps.put(second, currList);
			}
			steps.add(first);
			steps.add(second);
		}

		final int overhead = 1;//60;
		System.out.println(childParentMaps);
		final List<Character> finalOrder = new ArrayList<>();
		int time = 0;
		while (childParentMaps.size() > 0) {
			final List<Character> tmp = new ArrayList<>(steps);
			tmp.removeAll(childParentMaps.keySet());
			Collections.sort(tmp);
			System.out.println(tmp);
			if (tmp.size() > 0) {
				Character toRemove = tmp.get(0);
				int timeToAdd = (overhead + toRemove- 'A');
				System.out.println("timeToAdd " + timeToAdd);
				time += timeToAdd ;

				finalOrder.add(toRemove);
				final Map<Character, List<Character>> newChildParentMaps = new HashMap<>();
				for (final Map.Entry<Character, List<Character>> entry : childParentMaps.entrySet()) {
					final List<Character> newList = new ArrayList<>(entry.getValue());
					newList.remove(toRemove);
					if (newList.size() > 0) {
						newChildParentMaps.put(entry.getKey(), newList);
					}
				}
				childParentMaps = newChildParentMaps;
				steps.remove(toRemove);
			}
		}
		if (steps.size() == 1) {
			finalOrder.addAll(steps);
		}
		System.out.println(time);
		finalOrder.forEach(System.out::print);

	}

}
