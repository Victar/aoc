package adventofcode.year2021;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day14 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day14/input.txt");
		String init = "";
		final Map<String, String> rules = new HashMap<String, String>();
		for (final String input : data) {
			if (StringUtils.isNotEmpty(input)) {
				if (input.contains("->")) {
					final String[] arr = StringUtils.split(input, " -> ");
					rules.put(arr[0], arr[1]);
				} else {
					init = input;
				}
			}
		}
		System.out.println(init);

		for (int i = 0; i < 10; i++) {
			init = round(init, rules);
		}

		final Map<String, Integer> map = new HashMap<>();
		for (int i = 0; i < init.length(); i++) {
			final String current = "" + init.charAt(i);
			if (map.containsKey(current)) {
				map.put(current, map.get(current) + 1);
			} else {
				map.put(current, 1);
			}
		}
		long min = Long.MAX_VALUE;
		long max = Long.MIN_VALUE;

		for (final Map.Entry<String, Integer> entry : map.entrySet()) {
			min = Math.min(entry.getValue(), min);
			max = Math.max(entry.getValue(), max);

		}
		System.out.println(max - min);

	}

	private String round(final String input, final Map<String, String> rules) {
		char prev = input.charAt(0);
		final StringBuilder sb = new StringBuilder(prev);
		sb.append(prev);
		for (int i = 1; i < input.length(); i++) {
			final char current = input.charAt(i);
			final String key = prev + "" + current;
			if (rules.containsKey(key)) {
				sb.append(rules.get(key));
			}
			sb.append(current);
			prev = current;
		}
		return sb.toString();
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day14/input.txt");
		String init = "";
		final Map<String, String> rules = new HashMap<String, String>();
		for (final String input : data) {
			if (StringUtils.isNotEmpty(input)) {
				if (input.contains("->")) {
					final String[] arr = StringUtils.split(input, " -> ");
					rules.put(arr[0], arr[1]);
				} else {
					init = input;
				}
			}
		}

		Map<String, Long> mapRuleCount = stringToRuleCount(init);
		System.out.println(mapRuleCount);

		for (int i = 0; i < 40; i++) {
			mapRuleCount = doGoldRound(mapRuleCount, rules);
			System.out.println(mapRuleCount);

		}

		//Count results
		final Map<String, Long> map = new HashMap<>();
		map.put("" + init.charAt(init.length() - 1), 1l);

		for (final Map.Entry<String, Long> entry : mapRuleCount.entrySet()) {
			final String key = entry.getKey();
			final long count = entry.getValue();
			final String key0 = key.charAt(0) + "";
			if (map.containsKey(key0)) {
				map.put(key0, map.get(key0) + count);
			} else {
				map.put(key0, count);
			}
		}
		long min = Long.MAX_VALUE;
		long max = Long.MIN_VALUE;

		for (final Map.Entry<String, Long> entry : map.entrySet()) {
			min = Math.min(entry.getValue(), min);
			max = Math.max(entry.getValue(), max);

		}
		System.out.println(max - min);
	}

	private Map<String, Long> stringToRuleCount(final String input) {
		final Map<String, Long> map = new HashMap<>();
		char prev = input.charAt(0);
		for (int i = 1; i < input.length(); i++) {
			final char current = input.charAt(i);
			final String key = prev + "" + current;
			if (map.containsKey(key)) {
				map.put(key, map.get(key) + 1);
			} else {
				map.put(key, 1l);
			}
			prev = current;
		}
		return map;
	}

	private Map<String, Long> doGoldRound(final Map<String, Long> currentMap, final Map<String, String> rules) {
		final Map<String, Long> map = new HashMap<>();
		for (final Map.Entry<String, Long> entry : currentMap.entrySet()) {
			final String key = entry.getKey();
			final Long count = entry.getValue();
			if (rules.containsKey(key)) {
				final String newKey1 = key.charAt(0) + rules.get(key);
				final String newKey2 = rules.get(key) + key.charAt(1);
				if (map.containsKey(newKey1)) {
					map.put(newKey1, map.get(newKey1) + count);
				} else {
					map.put(newKey1, count);
				}
				if (map.containsKey(newKey2)) {
					map.put(newKey2, map.get(newKey2) + count);
				} else {
					map.put(newKey2, count);
				}
			}
		}
		return map;

	}
}
