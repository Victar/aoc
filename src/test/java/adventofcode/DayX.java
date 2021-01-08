package adventofcode;

import java.util.ArrayList;

import org.junit.Test;

public class DayX extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/dayX/input_sample.txt");
		int count = 0;
		for (String input : data) {
			count = input.length();
		}
		System.out.println(count);
	}

}
