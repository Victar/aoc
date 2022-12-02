package adventofcode.year2020;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Set;

import org.junit.Test;

import adventofcode.BaseTest;

public class Day14 extends BaseTest {

	@Test public void singleCheck() {
		final Set<Long> goldMaskSet = new HashSet<>();
		fillMaskSet("000000000000000000000000000000X1101X", goldMaskSet);
		System.out.println(goldMaskSet);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day14/input.txt");
		String mask = "";
		final Map<Long, Long> memory = new LinkedHashMap();
		for (final String input : data) {
			if (input.startsWith("mask")) {
				mask = input.substring(7);
			}
			if (input.startsWith("mem")) {
				final String value = input.substring(input.indexOf("=") + 1).trim();
				final Long valueL = Long.parseLong(value);
				final Integer key = Integer.parseInt(input.substring(4, input.indexOf("]")));
				final String goldMask = applyGoldMask(mask, key);
				final Set<Long> goldMaskSet = new HashSet<>();
				fillMaskSet(goldMask, goldMaskSet);
				for (final Long setValue : goldMaskSet) {
					memory.put(setValue, valueL);
				}
			}
		}
		long result = 0;
		for (final Map.Entry<Long, Long> entry : memory.entrySet()) {
			result += entry.getValue();
		}
		System.out.println(result);
	}

	public void fillMaskSet(final String mask, final Set<Long> longValue) {
		if (mask.contains("X")) {
			fillMaskSet(mask.replaceFirst("X", "0"), longValue);
			fillMaskSet(mask.replaceFirst("X", "1"), longValue);
		} else {
			final long value = new BigInteger(mask, 2).longValue();
			longValue.add(value);
		}
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day14/input.txt");
		String mask = "";
		final Map<Integer, Long> keyValue = new LinkedHashMap();
		for (final String input : data) {
			if (input.startsWith("mask")) {
				mask = input.substring(7);
			}
			if (input.startsWith("mem")) {
				final String value = input.substring(input.indexOf("=") + 1).trim();
				final Integer key = Integer.parseInt(input.substring(4, input.indexOf("]")));
				final Long valueL = Long.parseLong(value);
				keyValue.put(key, applyMask(mask, valueL));
			}
		}

		long result = 0;
		for (final Map.Entry<Integer, Long> entry : keyValue.entrySet()) {
			result += entry.getValue().longValue();
		}
		System.out.println(result);
	}

	public String applyGoldMask(final String mask, final long value) {
		final String valueStr = Long.toBinaryString(value);
		final String zeroMask = "000000000000000000000000000000000000";
		final String finalValue = zeroMask.substring(valueStr.length()) + valueStr;
		final StringBuilder sb = new StringBuilder();
		for (int i = 0; i < finalValue.length(); i++) {
			final char maskChar = mask.charAt(i);
			if ("0".equalsIgnoreCase(maskChar + "")) {
				sb.append(finalValue.charAt(i));
			} else {
				sb.append(maskChar);
			}
		}
		return sb.toString();
	}

	public long applyMask(final String mask, final long value) {
		final String valueStr = Long.toBinaryString(value);
		final String zeroMask = "000000000000000000000000000000000000";
		final String finalValue = zeroMask.substring(valueStr.length()) + valueStr;
		final StringBuilder sb = new StringBuilder();
		for (int i = 0; i < finalValue.length(); i++) {
			final char maskChar = mask.charAt(i);
			if ("X".equalsIgnoreCase(maskChar + "")) {
				sb.append(finalValue.charAt(i));
			} else {
				sb.append(maskChar);
			}
		}
		return new BigInteger(sb.toString(), 2).longValue();
	}

}

