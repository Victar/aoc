package adventofcode.year2020;

import java.util.ArrayList;

import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day6 extends BaseTest {

	@Test @Ignore public void singleCheck() {
		System.out.println(getInterceptionStr("asd", "abmd"));
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day6/input.txt");
		int count = 0;
		String currentGroups = "";
		for (final String input : data) {
			if (StringUtils.isEmpty(input)) {
				final int current = uniqueCharacters(currentGroups);
				count += current;
				currentGroups = "";

			} else {
				currentGroups += input;
			}
		}
		final int current = uniqueCharacters(currentGroups);
		count += current;
		System.out.println(count);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day6/input.txt");
		int count = 0;
		String currentGroups = "";
		boolean isEmpty = true;
		for (final String input : data) {
			if (StringUtils.isEmpty(input)) {
				final int current = uniqueCharacters(currentGroups);
				count += current;
				currentGroups = "";
				isEmpty = true;

			} else {
				if (StringUtils.isEmpty(currentGroups) && isEmpty) {
					currentGroups = input;
				} else {
					isEmpty = false;
					currentGroups = getInterceptionStr(currentGroups, input);
				}
			}
		}
		final int current = uniqueCharacters(currentGroups);
		count += current;
		System.out.println(count);
	}

}
