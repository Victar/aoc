package adventofcode;

import java.util.ArrayList;

import org.junit.Test;

public class DayX extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/dayX/input_sample.txt");
		for (final String input : data) {
			System.out.println(input);
		}
	}

}
